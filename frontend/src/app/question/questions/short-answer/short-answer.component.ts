import {
  Component,
  EventEmitter,
  Input,
  OnChanges,
  OnInit,
  Output,
  SimpleChanges,
} from '@angular/core';
import { Question } from '../../question';

@Component({
  selector: 'app-short-answer',
  templateUrl: './short-answer.component.html',
  styleUrls: ['./short-answer.component.scss'],
})
export class ShortAnswerComponent implements OnInit, OnChanges {
  @Input() question: Question;
  @Output() answeredQuestion = new EventEmitter<Question>();
  answer: string;
  private questionCopy: Question;

  constructor() {}

  ngOnInit(): void {
    // b/c you can't change an @Input I'm copying
    this.questionCopy = { ...this.question };
  }

  onQuestionChanged(value: string): void {
    this.questionCopy.selectedChoice = value;
    this.answeredQuestion.emit(this.questionCopy);
  }

  ngOnChanges(changes: SimpleChanges): void {
    this.questionCopy.answer = null;
  }
}
