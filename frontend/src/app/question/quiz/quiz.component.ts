import { Component, HostListener, OnInit } from '@angular/core';
import { timer } from 'rxjs';
import { Question } from '../question';
import { QuestionService } from '../question.service';
import { QuestionType } from '../question-type.enum';

@Component({
  selector: 'app-quiz',
  templateUrl: './quiz.component.html',
  styleUrls: ['./quiz.component.scss'],
})
export class QuizComponent implements OnInit {
  // timer vars
  public time = 0;
  private interval;
  private subscribeTimer: any;

  // question vars
  public questionIndex = 0;
  public questions: Array<Question>;

  // making enum available to html
  public allQuestionTypes = QuestionType;
  public canAdvance: boolean;

  constructor(private questionService: QuestionService) {}

  ngOnInit(): void {
    this.canAdvance = false;
    const source = timer(1000, 2000);
    source.subscribe((value) => {
      this.subscribeTimer = this.time - value;
    });
    this.startTimer();
    this.questionService.getFiveQuestions();
    this.questionService.questions.subscribe((data) => (this.questions = data));
  }

  answerChanged(question: Question): void {
    this.questions[this.questionIndex].selectedChoice = question.selectedChoice;
    this.canAdvance = question.selectedChoice != null;
  }

  private startTimer(): void {
    this.interval = setInterval(() => {
      this.time++;
    }, 1000);
  }

  getReadableTime(time: number): string {
    const min = Math.floor(this.time / 60);
    const sec =
      this.time % 60 < 10 ? `0${this.time % 60}` : `${this.time % 60}`;
    return `${min}:${sec}`;
  }

  private pauseTimer(): void {
    clearInterval(this.interval);
  }

  public getSelectedQuestionType(): QuestionType {
    return this.questions[this.questionIndex].type;
  }

  /*@HostListener('window:beforeunload', ['$event']) unloadHandler(
    event: Event
  ): void {
    const result = confirm('Changes you made may not be saved.');
    if (result) {
      // Do more processing...
    }
    event.returnValue = false; // stay on same page
  }*/

  advanceQuestion(): void {
    if (this.questionIndex === 4) {
      // navigate to results
    }
    this.canAdvance = false;
    this.questionIndex++;
  }
}
