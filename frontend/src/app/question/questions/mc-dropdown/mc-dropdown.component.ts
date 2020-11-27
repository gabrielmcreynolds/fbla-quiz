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
  constructor() {}

  ngOnInit(): void {}

  onChange(value: string): void {
    this.question.selectedChoice = value;
    this.answeredQuestion.emit(this.question);
  }
}
