import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './auth/login/login.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { QuizComponent } from './question/quiz/quiz.component';
import { ResultsComponent } from './results/results.component';

const routes: Routes = [
  { path: '', component: LoginComponent },
  { path: 'dashboard', component: DashboardComponent },
  { path: 'quiz', component: QuizComponent },
  { path: 'results', component: ResultsComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { relativeLinkResolution: 'legacy' })],
  exports: [RouterModule],
  // providers: [AuthGuard]
})
export class AppRoutingModule {}
