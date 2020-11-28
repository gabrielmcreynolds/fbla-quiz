import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Question } from '../../question';

@Component({
  selector: 'app-mc-dropdown',
  templateUrl: './mc-dropdown.component.html',
  styleUrls: ['./mc-dropdown.component.scss'],
})
export class McDropdownComponent implements OnInit {
  @Input() question: Question;
  @Output() answeredQuestion = new EventEmitter<Question>();
  selectedChoice: string;
  private questionCopy: Question;
  constructor() {}

  ngOnInit(): void {
    // b/c you can't change an @Input I'm copying
    this.questionCopy = { ...this.question };
  }

  onChange(value: string): void {
    this.questionCopy.selectedChoice = value;
    this.answeredQuestion.emit(this.questionCopy);
  }
}
