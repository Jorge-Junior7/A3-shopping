import { Component } from '@angular/core';
import { RouterLink } from '@angular/router';
import { LoginComponent } from '../login/login.component';
import { RegisterService } from '../../services/api/register.service';


@Component({
  selector: 'app-final-registration',
  standalone: true,
  imports: [RouterLink, LoginComponent, ],
  templateUrl: './final-registration.component.html',
  styleUrl: './final-registration.component.css'
})
export class FinalRegistrationComponent {

}
