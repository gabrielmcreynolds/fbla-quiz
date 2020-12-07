import { Component, OnInit } from '@angular/core';
import {
  AbstractControl,
  FormControl,
  FormGroup,
  ValidationErrors,
  ValidatorFn,
  Validators,
} from '@angular/forms';
import { ValidateFn } from 'codelyzer/walkerFactory/walkerFn';

@Component({
  selector: 'app-create-account',
  templateUrl: './create-account.component.html',
  styleUrls: ['./create-account.component.scss'],
})
export class CreateAccountComponent implements OnInit {
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

  constructor() {}

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

  validate(): void {
    if (this.form.valid) {
      console.log('Valid');
    } else {
      console.log('Invalid');
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
