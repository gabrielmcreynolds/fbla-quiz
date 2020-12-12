import { Component, OnInit } from '@angular/core';
import {
  AbstractControl,
  FormControl,
  FormGroup,
  ValidationErrors,
  ValidatorFn,
  Validators,
} from '@angular/forms';
import { AuthService } from '../auth.service';
import { AuthStatus } from '../auth-status.enum';
import { Router } from '@angular/router';

@Component({
  selector: 'app-create-account',
  templateUrl: './create-account.component.html',
  styleUrls: ['./create-account.component.scss'],
})
export class CreateAccountComponent implements OnInit {
  incorrectEmail = false;

  get name(): AbstractControl {
    return this.form.get('name');
  }

  get email(): AbstractControl {
    return this.form.get('email');
  }

  get passcode(): AbstractControl {
    return this.form.get('passcode');
  }

  get confirmPasscode(): AbstractControl {
    return this.form.get('confirmPasscode');
  }

  constructor(private authService: AuthService, private router: Router) {}

  form = new FormGroup(
    {
      name: new FormControl('', [Validators.required]),
      email: new FormControl('', [Validators.required, Validators.email]),
      passcode: new FormControl('', [Validators.required]),
      confirmPasscode: new FormControl('', [Validators.required]),
    },
    { validators: passwordValidator }
  );

  ngOnInit(): void {}

  createAccount(): void {
    if (this.form.valid) {
      this.authService
        .createUser({
          name: this.name.value,
          email: this.email.value,
          password: this.passcode.value,
        })
        .subscribe((authStatus) => {
          if (authStatus === AuthStatus.LoggedIn) {
            this.router.navigate(['/dashboard']);
          }
          this.incorrectEmail = authStatus === AuthStatus.IncorrectEmail;
        });
    }
  }
}

export const passwordValidator: ValidatorFn = (
  control: FormGroup
): ValidationErrors | null => {
  const passcode = control.get('passcode');
  const confPasscode = control.get('confirmPasscode');
  return passcode && confPasscode && passcode?.value === confPasscode?.value
    ? null
    : { samePassword: true };
};
