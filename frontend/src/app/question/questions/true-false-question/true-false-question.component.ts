import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
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
export class TrueFalseQuestionComponent implements OnInit {
  @Input() question: Question;
  @Output() answeredQuestion = new EventEmitter<Question>();
  selectedAnswer = true;

  constructor() {}

  ngOnInit(): void {
    this.question.selectedChoice = this.selectedAnswer;
    this.answeredQuestion.emit(this.question);
  }

  setAnswer(answer: boolean): void {
    this.selectedAnswer = answer;
    this.question.selectedChoice = this.selectedAnswer;
    this.answeredQuestion.emit(this.question);
  }
}
