import { Component } from '@angular/core';
import { RouterLink, RouterOutlet } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';
import { HomeComponent } from './pages/home/home.component';
import { MenuComponent } from './components/menu/menu.component';
import { MyAdsComponent } from './pages/my-ads/my-ads.component';
import { LoginResetComponent } from './pages/login-reset/login-reset.component';
import { RegisterComponent } from './pages/register/register.component';
import { ProductAddComponent } from './pages/product-add/product-add.component';
import { FinalRegistrationComponent } from './pages/final-registration/final-registration.component';
import { MyInformationComponent } from './pages/my-information/my-information.component';
import { ProfileComponent } from './pages/profile/profile.component';
import { SavedComponent } from './pages/saved/saved.component';
import { ChatComponent } from './pages/chat/chat.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, RouterLink, LoginComponent, HomeComponent, MenuComponent, MyAdsComponent, LoginResetComponent, RegisterComponent, ProductAddComponent, FinalRegistrationComponent, MyInformationComponent, ProfileComponent, SavedComponent, ChatComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'shopping';
}
