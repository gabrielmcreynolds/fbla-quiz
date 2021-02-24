import { Component, OnInit } from '@angular/core';
import {
  AbstractControl,
  FormControl,
  FormGroup,
  Validators,
} from '@angular/forms';
import { AuthService } from '../auth.service';
import { AuthStatus } from '../auth-status.enum';

@Component({
  selector: 'app-home',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
})
export class LoginComponent implements OnInit {
  emailError = false;
  passwordError = false;

  loginForm = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [Validators.required]),
  });

  constructor(private authService: AuthService) {}

  get email(): AbstractControl {
    return this.loginForm.get('email');
  }

  get password(): AbstractControl {
    return this.loginForm.get('password');
  }

  login(): void {
    if (this.loginForm.valid) {
      this.authService
        .login(this.email.value, this.password.value)
        .subscribe((authStatus) => {
          if (authStatus === AuthStatus.IncorrectPassword) {
            this.emailError = false;
            this.passwordError = true;
          } else if (authStatus === AuthStatus.IncorrectEmail) {
            this.emailError = true;
            this.passwordError = false;
          }
        });
    }
  }

  ngOnInit(): void {}
}
