import {Component, OnInit} from '@angular/core';
import {Question} from '../questions/question';
import {HttpClient} from '@angular/common/http';
import {AuthService} from '../auth/auth.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {
  questions: Array<Question>;

  constructor(private http: HttpClient, private authService: AuthService) {
  }

  ngOnInit(): void {
  }

  getQuestions(): void {
    this.authService.logout();
  }

}
