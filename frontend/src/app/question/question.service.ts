import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Question } from './question';
import { BehaviorSubject } from 'rxjs';
import { QuestionType } from './question-type.enum';
import { AuthService } from '../auth/auth.service';
import { User } from '../auth/user';

@Injectable({
  providedIn: 'root',
})
export class QuestionService {
  constructor(private http: HttpClient, private authService: AuthService) {}
  time: number;
  correctQuestions: number;

  questions = new BehaviorSubject<Array<Question>>(null);

  private static determineQuestionType(question: Question): QuestionType {
    if (typeof question.correctChoice === 'boolean') {
      return QuestionType.TrueFalse;
    }
    if (question.choices == null) {
      return QuestionType.ShortAnswer;
    }
    return QuestionService.randomMC();
  }

  private static randomMC(): QuestionType {
    const rand = Math.floor(Math.random() * Math.floor(2));
    if (rand === 0) {
      return QuestionType.McDropdown;
    } else {
      return QuestionType.McPickBox;
    }
  }

  private static isQuestionCorrect(question: Question): boolean {
    let result: boolean;
    if (typeof question.correctChoice === 'boolean') {
      result = question.correctChoice === question.selectedChoice;
    } else {
      result =
        question.selectedChoice.toString().trim().toLowerCase() ===
        question.correctChoice.toString().trim().toLowerCase();
    }
    return result;
  }

  getFiveQuestions(): void {
    this.http
      .get<{ questions: Array<Question> }>('questions')
      .subscribe((data) => {
        for (const question of data.questions) {
          question.type = QuestionService.determineQuestionType(question);
        }
        this.questions.next(data.questions);
      });
  }

  setQuestions(questions: Array<Question>, time: number): void {
    let correctQuestions = 0;
    questions.forEach((question) => {
      if (QuestionService.isQuestionCorrect(question)) {
        correctQuestions++;
        question.isCorrect = true;
      } else {
        question.isCorrect = false;
      }
    });
    this.correctQuestions = correctQuestions;
    this.time = time;
    this.questions.next(questions);
    this.http
      .post<{ message: string; user: User }>('users/addTest', {
        score: correctQuestions,
        time,
      })
      .subscribe((res) => {
        if (res) {
          this.authService.setUser(res.user);
        }
      });
  }
}
