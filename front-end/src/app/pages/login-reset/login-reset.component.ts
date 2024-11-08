import { NgIf } from '@angular/common';
import { Component } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';

@Component({
  selector: 'app-login-reset',
  standalone: true,
  imports: [ReactiveFormsModule, NgIf],
  templateUrl: './login-reset.component.html',
  styleUrls: ['./login-reset.component.css']
})
export class LoginResetComponent {
  resetForm: FormGroup;
  changePasswordForm: FormGroup;
  errorMessage: string | null = null;
  isVerified: boolean = false; // Controla se os dados iniciais foram confirmados

  constructor(private fb: FormBuilder) {
    // Formulário inicial para verificação dos dados
    this.resetForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      cpf: ['', [Validators.required, Validators.pattern(/^\d{3}\.\d{3}\.\d{3}-\d{2}$/)]],
      birthdate: ['', Validators.required],
      recoveryPhrase: ['', Validators.required]
    });

    // Formulário para alteração da senha
    this.changePasswordForm = this.fb.group({
      newPassword: ['', [
        Validators.required,
        Validators.pattern(/^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#\$%\^&\*])(?=.{8,})/)
      ]],
      confirmPassword: ['', Validators.required]
    }, { validators: this.passwordMatchValidator });
  }

  // Validador para verificar se as senhas correspondem
  passwordMatchValidator(form: FormGroup) {
    const password = form.get('newPassword')?.value;
    const confirmPassword = form.get('confirmPassword')?.value;
    return password === confirmPassword ? null : { mismatch: true };
  }

  // Envia os dados iniciais para validação
  verifyData(): void {
    if (this.resetForm.valid) {
      const formData = this.resetForm.value;
      console.log('Dados para verificação:', formData);

      // Suponha que chamamos uma API para validar os dados
      // Se os dados forem válidos, definimos isVerified como true
      this.isVerified = true; // A API confirmou os dados
      this.errorMessage = null; // Limpa qualquer mensagem de erro
    } else {
      this.errorMessage = 'Por favor, preencha todos os campos corretamente.';
    }
  }

  // Atualiza a senha após a confirmação
  updatePassword(): void {
    if (this.changePasswordForm.valid) {
      const newPassword = this.changePasswordForm.get('newPassword')?.value;
      console.log('Nova senha:', newPassword);

      // Aqui, você chamaria uma API para atualizar a senha no backend
      alert('Senha alterada com sucesso!');
    } else {
      this.errorMessage = 'Certifique-se de que as senhas correspondem e cumprem os requisitos.';
    }
  }
}
