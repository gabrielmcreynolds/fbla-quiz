import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { RouterModule } from '@angular/router';
import { LoginComponent } from './auth/login/login.component';
import { AppRoutingModule } from './app-routing.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { AuthInterceptor } from './auth/auth.interceptor';
import { DashboardComponent } from './dashboard/dashboard.component';
import { QuizComponent } from './question/quiz/quiz.component';
import { TrueFalseQuestionComponent } from './question/questions/true-false-question/true-false-question.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { McPickBoxComponent } from './question/questions/mc-pick-box/mc-pick-box.component';
import { ShortAnswerComponent } from './question/questions/short-answer/short-answer.component';
import { McDropdownComponent } from './question/questions/mc-dropdown/mc-dropdown.component';
import { SpinnerComponent } from './spinner/spinner.component';
import { ResultsComponent } from './results/results.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    DashboardComponent,
    QuizComponent,
    TrueFalseQuestionComponent,
    McPickBoxComponent,
    ShortAnswerComponent,
    McDropdownComponent,
    SpinnerComponent,
    ResultsComponent,
  ],
  imports: [
    BrowserAnimationsModule,
    BrowserModule,
    RouterModule,
    AppRoutingModule,
    ReactiveFormsModule,
    FormsModule,
    HttpClientModule,
  ],
  providers: [
    { provide: HTTP_INTERCEPTORS, useClass: AuthInterceptor, multi: true },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
