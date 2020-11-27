import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Question } from '../../question';

@Component({
  selector: 'app-short-answer',
  templateUrl: './short-answer.component.html',
  styleUrls: ['./short-answer.component.scss'],
})
export class ShortAnswerComponent implements OnInit {
  @Input() question: Question;
  @Output() answeredQuestion = new EventEmitter<Question>();
  answer: string;

  constructor() {}

  ngOnInit(): void {}

  onQuestionChanged(value: string): void {
    this.question.selectedChoice = value;
    this.answeredQuestion.emit(this.question);
  }
}
