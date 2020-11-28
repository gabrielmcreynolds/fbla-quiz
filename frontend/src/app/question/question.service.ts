import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Question } from './question';
import { BehaviorSubject } from 'rxjs';
import { QuestionType } from './question-type.enum';

@Injectable({
  providedIn: 'root',
})
export class QuestionService {
  constructor(private http: HttpClient) {}

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
}
