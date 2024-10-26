import { Component, OnInit } from '@angular/core';
import { Router, RouterLink, RouterOutlet } from '@angular/router';
import { AuthService } from '../../services/api/auth.service';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ReactiveFormsModule } from '@angular/forms';
import { NgClass, NgIf } from '@angular/common';
import { LoginResetComponent } from '../login-reset/login-reset.component';
import { RegisterComponent } from '../register/register.component';


@Component({
  selector: 'app-login',
  standalone: true,
  imports: [ReactiveFormsModule, NgIf, NgClass, LoginResetComponent, RegisterComponent, RouterOutlet, RouterLink],
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginForm: FormGroup;
  errorMessage: string = '';
  isLoading: boolean = false;

  constructor(
    private fb: FormBuilder,
    private authService: AuthService,
    private router: Router
  ) {
    this.loginForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required]]
    });
  }

  ngOnInit(): void {}

  onLogin(): void {
    if (this.loginForm.valid) {
      this.isLoading = true; // Começar o estado de carregamento
      const { email, password } = this.loginForm.value;

      this.authService.login(email, password).subscribe({
        next: (response) => {
          console.log('Login bem-sucedido', response);
          this.router.navigate(['/home']);
        },
        error: (error) => {
          this.isLoading = false; // Parar o estado de carregamento
          console.error('Erro no login', error);
          this.errorMessage = this.getErrorMessage(error);
        }
      });
    } else {
      this.errorMessage = 'Por favor, pré-encha os campos corretamente.';
    }
  }

  private getErrorMessage(error: any): string {
    // Ajustar a mensagem de erro com base na resposta da API
    if (error.error && error.error.message) {
      return error.error.message; // Mensagem específica da API
    }
    return 'Erro desconhecido. Tente novamente mais tarde.';
  }
}
