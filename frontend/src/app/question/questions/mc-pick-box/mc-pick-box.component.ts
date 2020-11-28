import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Question } from '../../question';
import {
  animate,
  state,
  style,
  transition,
  trigger,
} from '@angular/animations';

@Component({
  selector: 'app-mc-pick-box',
  templateUrl: './mc-pick-box.component.html',
  styleUrls: ['./mc-pick-box.component.scss'],
  animations: [
    trigger('mc', [
      state(
        'selected',
        style({
          backgroundColor: '#22D1EE',
        })
      ),

      state(
        'not-selected',
        style({
          backgroundColor: '#B8E7EC',
        })
      ),

      transition('* => *', [animate('.3s')]),
    ]),
  ],
})
export class McPickBoxComponent implements OnInit {
  @Input() question: Question;
  @Output() answeredQuestion = new EventEmitter<Question>();
  selectedChoice: string;
  private questionCopy: Question;

  constructor() {}

  ngOnInit(): void {
    // b/c you can't change an @Input I'm copying
    this.questionCopy = { ...this.question };
  }

  setChoice(choice: string): void {
    this.selectedChoice = choice;
    this.questionCopy.selectedChoice = this.selectedChoice;
    this.answeredQuestion.emit(this.questionCopy);
  }
}
