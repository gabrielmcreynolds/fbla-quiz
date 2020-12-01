import { Component, OnInit } from '@angular/core';
import { QuestionService } from '../question/question.service';
import { Question } from '../question/question';

@Component({
  selector: 'app-results',
  templateUrl: './results.component.html',
  styleUrls: ['./results.component.scss'],
})
export class ResultsComponent implements OnInit {
  public questions: Array<Question>;
  public correctQuestions: number;

  constructor(private questionService: QuestionService) {}

  ngOnInit(): void {
    this.questions = this.questionService.questions.value;
    this.correctQuestions = this.questionService.correctQuestions;
  }

  getReadableTime(): string {
    const min = Math.floor(this.questionService.time / 60);
    const sec =
      this.questionService.time % 60 < 10
        ? `0${this.questionService.time % 60}`
        : `${this.questionService.time % 60}`;
    return `${min}:${sec}`;
  }
}
