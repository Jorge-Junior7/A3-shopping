import { Component, ViewEncapsulation } from '@angular/core';
import { RouterLink, RouterOutlet } from '@angular/router';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { RegisterService } from '../../services/api/register.service';
import { CommonModule, NgIf } from '@angular/common';
import { FinalRegistrationComponent } from '../final-registration/final-registration.component';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [RouterLink, RouterOutlet, CommonModule, ReactiveFormsModule, NgIf, FinalRegistrationComponent],
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css'],
  encapsulation: ViewEncapsulation.None,
})

export class RegisterComponent {
  registerForm: FormGroup;
  errorMessage: string | null = null;
  selectedFile: File | null = null;

  constructor(private fb: FormBuilder, private registerService: RegisterService) {
    this.registerForm = this.fb.group({
      full_name: ['', Validators.required],
      birthdate: ['', Validators.required],
      cpf: ['', Validators.required],
      nickname: [''],
      location: [''],
      email: ['', [Validators.required, Validators.email]],
      password: ['', [
        Validators.required,
        Validators.pattern(/^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#\$%\^&\*])(?=.{8,})/)
      ]]
    });
  }

  onFileSelected(event: Event): void {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files.length > 0) {
      this.selectedFile = input.files[0];
      const reader = new FileReader();
      reader.onload = (e: any) => {
        const img = document.getElementById('profilePhoto') as HTMLImageElement;
        img.src = e.target.result;
        img.style.display = 'block';
      };
      reader.readAsDataURL(this.selectedFile);
    }
  }

  submitForm(): void {
    console.log('Dados do formulário:', this.registerForm.value);

    if (this.registerForm.valid) {
      const formData = new FormData();
      Object.keys(this.registerForm.controls).forEach(key => {
        formData.append(key, this.registerForm.controls[key].value);
      });

      if (this.selectedFile) {
        formData.append('profilePhoto', this.selectedFile);
      }

      console.log('Dados a serem enviados:', formData);

      this.registerService.register(formData).subscribe({
        next: (response) => {
          console.log('Usuário registrado com sucesso!', response);
        },
        error: (error) => {
          this.errorMessage = error.error.errors ? Object.values(error.error.errors).join(', ') : 'Erro ao registrar';
          console.error('Erro ao registrar usuário:', error);
        }
      });
    } else {
      console.error('Formulário inválido:', this.registerForm.errors);
      this.errorMessage = 'Preencha todos os campos obrigatórios corretamente.';

      for (const control in this.registerForm.controls) {
        if (this.registerForm.controls[control].invalid) {
          const errors = this.registerForm.controls[control].errors;
          console.log(`Campo ${control} inválido:`, errors);

          if (errors?.['required']) {
            this.errorMessage += ` ${control} é obrigatório.`;
          }
          if (errors?.['pattern'] && control === 'password') {
            this.errorMessage += ` A senha deve ter pelo menos 8 caracteres, uma letra maiúscula, uma letra minúscula, um número e um caractere especial.`;
          }
        }
      }
    }
  }
}
