import {Component, OnInit} from '@angular/core';
import {Question} from '../questions/question';
import {HttpClient} from '@angular/common/http';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {
  questions: Array<Question>;

  constructor(private http: HttpClient) {
  }

  ngOnInit(): void {
  }

  getQuestions(): void {
    this.http.get<{ questions: Array<Question> }>('questions')
      .subscribe(value => {
        console.log('Question: ');
        console.log(value);
        this.questions = value.questions;
      });
  }

}
