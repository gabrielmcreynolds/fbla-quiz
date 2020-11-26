import { Component, OnInit } from '@angular/core';
import { Question } from '../question/question';
import { HttpClient } from '@angular/common/http';
import { AuthService } from '../auth/auth.service';
import { User } from '../auth/user';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss'],
})
export class DashboardComponent implements OnInit {
  questions: Array<Question>;
  user: User;
  isLoading = true;

  constructor(
    private router: Router,
    private http: HttpClient,
    private authService: AuthService
  ) {}

  ngOnInit(): void {
    this.authService.getUser().subscribe((user) => {
      this.user = user;
      this.isLoading = user == null;
    });
  }

  getAverageTime(): string {
    const totalSeconds = this.user.totalTime / this.user.testsTaken;
    const minutes = Math.floor(totalSeconds / 60);
    const seconds = totalSeconds % 60;
    if (minutes === 0) {
      return `${seconds} s`;
    }
    if (seconds === 0) {
      return `${minutes} m`;
    }
    return `${minutes} m ${seconds} s`;
  }

  startQuiz(): void {
    this.router.navigate(['/quiz']);
  }
}
