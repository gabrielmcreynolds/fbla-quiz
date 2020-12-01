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
import {
  trigger,
  state,
  style,
  animate,
  transition,
} from '@angular/animations';

@Component({
  selector: 'app-true-false-question',
  templateUrl: './true-false-question.component.html',
  styleUrls: ['./true-false-question.component.scss'],
  animations: [
    trigger('trueFalse', [
      state(
        'selected',
        style({
          color: 'white',
          backgroundColor: '#3D5AF1',
        })
      ),

      state(
        'not-selected',
        style({
          color: '#22D1EE',
          backgroundColor: 'transparent',
        })
      ),

      transition('* => *', [animate('.3s')]),
    ]),
  ],
})
export class TrueFalseQuestionComponent implements OnInit, OnChanges {
  @Input() question: Question;
  @Output() answeredQuestion = new EventEmitter<Question>();
  selectedAnswer: boolean;
  private questionCopy: Question;

  constructor() {}

  ngOnInit(): void {
    // b/c you can't change an @Input I'm copying
    this.questionCopy = { ...this.question };
  }

  setAnswer(answer: boolean): void {
    this.selectedAnswer = answer;
    this.questionCopy.selectedChoice = this.selectedAnswer;
    this.answeredQuestion.emit(this.questionCopy);
  }

  ngOnChanges(changes: SimpleChanges): void {
    this.selectedAnswer = null;
  }
}
