import { Component, HostListener, OnInit } from '@angular/core';
import { timer } from 'rxjs';
import { Question } from '../question';

@Component({
  selector: 'app-quiz',
  templateUrl: './quiz.component.html',
  styleUrls: ['./quiz.component.scss'],
})
export class QuizComponent implements OnInit {
  public seconds: string;
  public minutes: number;
  private time = 0;
  private interval;
  private subscribeTimer: any;
  public question: Question;

  constructor() {}

  ngOnInit(): void {
    const source = timer(1000, 2000);
    source.subscribe((value) => {
      this.subscribeTimer = this.time - value;
    });
    this.startTimer();

    this.question = {
      question: 'Was FBLA invented in 1786',
      correctChoice: '1785',
      choices: ['1785', '1786', '1787', '1788'],
    };
  }

  startTimer(): void {
    this.interval = setInterval(() => {
      this.time++;
      this.minutes = Math.floor(this.time / 60);
      this.seconds =
        this.time % 60 < 10 ? `0${this.time % 60}` : `${this.time % 60}`;
    }, 1000);
  }

  pauseTimer(): void {
    clearInterval(this.interval);
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
}
