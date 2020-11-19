import {Component, OnInit} from '@angular/core';
import {Question} from '../questions/question';
import {HttpClient} from '@angular/common/http';
import {AuthService} from '../auth/auth.service';
import {User} from '../auth/user';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {
  questions: Array<Question>;
  user: User;

  constructor(private http: HttpClient, private authService: AuthService) {
  }

  ngOnInit(): void {
    this.authService.getUser().subscribe(user => this.user = user);
  }

  getQuestions(): void {
    this.authService.logout();
  }

}
