import { NgIf } from '@angular/common';
import { Component } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { VerifyService } from '../../services/api/verify.service';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-login-reset',
  standalone: true,
  imports: [ReactiveFormsModule, NgIf, RouterLink],
  templateUrl: './login-reset.component.html',
  styleUrls: ['./login-reset.component.css']
})
export class LoginResetComponent {
  resetForm: FormGroup;
  changePasswordForm: FormGroup;
  errorMessage: string | null = null;
  isVerified: boolean = false;
  passwordUpdated: boolean = false;

  constructor(private fb: FormBuilder, private verifyService: VerifyService) {
    this.resetForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      cpf: ['', [Validators.required, Validators.pattern(/^\d{11}$/)]],
      birthdate: ['', Validators.required],
      recoveryPhrase: ['', Validators.required]
    });

    this.changePasswordForm = this.fb.group({
      newPassword: ['', [
        Validators.required,
        Validators.pattern(/^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#\\$%\\^&\\*])(?=.{8,})/)
      ]],
      confirmPassword: ['', Validators.required]
    }, { validators: this.passwordMatchValidator });
  }

  passwordMatchValidator(form: FormGroup) {
    const password = form.get('newPassword')?.value;
    const confirmPassword = form.get('confirmPassword')?.value;
    return password === confirmPassword ? null : { mismatch: true };
  }

  verifyData(): void {
    if (this.resetForm.valid) {
      const formData = {
        email: this.resetForm.get('email')?.value,
        cpf: this.resetForm.get('cpf')?.value,
        birth_date: new Date(this.resetForm.get('birthdate')?.value).toISOString().split('T')[0],
        recovery_phrase: this.resetForm.get('recoveryPhrase')?.value
      };

      this.verifyService.verifyUserData(formData).subscribe(
        () => {
          this.isVerified = true;
          this.errorMessage = null;
        },
        (error) => {
          this.errorMessage = error.error.message || 'Palavra-chave não reconhecida ou não registrada';
        }
      );
    } else {
      this.errorMessage = 'Por favor, preencha todos os campos corretamente.';
    }
  }

  updatePassword(): void {
    if (this.changePasswordForm.valid) {
      const formData = {
        email: this.resetForm.get('email')?.value,
        new_password: this.changePasswordForm.get('newPassword')?.value
      };

      this.verifyService.updateUserPassword(formData).subscribe(
        () => {
          this.passwordUpdated = true;
          this.isVerified = false;
          this.errorMessage = null;
        },
        (error) => {
          this.errorMessage = error.error.message || 'Erro ao atualizar a senha';
        }
      );
    } else {
      this.errorMessage = 'Certifique-se de que as senhas correspondem e cumprem os requisitos.';
    }
  }
}
