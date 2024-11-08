import { Component, ViewEncapsulation } from '@angular/core';
import { Router, RouterLink, RouterOutlet } from '@angular/router';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { RegisterService } from '../../services/api/register.service';
import { CommonModule, NgIf } from '@angular/common';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [RouterLink, RouterOutlet, CommonModule, ReactiveFormsModule, NgIf],
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css'],
  encapsulation: ViewEncapsulation.None,
})

export class RegisterComponent {
  registerForm: FormGroup;
  errorMessage: string | null = null;
  selectedFile: File | null = null;

  constructor(
    private fb: FormBuilder,
    private registerService: RegisterService,
    private router: Router
  ) {
    this.registerForm = this.fb.group({
      full_name: ['', Validators.required],
      birthdate: ['', Validators.required],
      cpf: ['', [Validators.required, Validators.pattern(/^\d{3}\.\d{3}\.\d{3}-\d{2}$/)]],
      nickname: ['', Validators.required],
      location: ['', Validators.required],
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
    this.errorMessage = null;

    if (this.registerForm.valid) {
      const formData = new FormData();
      Object.keys(this.registerForm.controls).forEach(key => {
        formData.append(key, this.registerForm.controls[key].value);
      });

      if (this.selectedFile) {
        formData.append('profilePhoto', this.selectedFile);
      }

      this.registerService.register(formData).subscribe({
        next: (response) => {
          const recoveryPhrase = response.recovery_phrase;  // Captura a frase de recuperação da resposta
          this.router.navigate(['/final-registration'], { queryParams: { recoveryPhrase } });  // Redireciona com a frase de recuperação
        },
        error: (error) => {
          this.errorMessage = error.error.errors ? Object.values(error.error.errors).join(', ') : 'Erro ao registrar';
          console.error('Erro ao registrar usuário:', error);
        }
      });
    } else {
      this.errorMessage = 'Preencha todos os campos obrigatórios corretamente.';
    }
  }
}
