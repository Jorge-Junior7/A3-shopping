import { Routes } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';
import { HomeComponent } from './pages/home/home.component';
import { MyAdsComponent } from './pages/my-ads/my-ads.component';
import { ProfileComponent } from './pages/profile/profile.component';
import { LoginResetComponent } from './pages/login-reset/login-reset.component';
import { RegisterComponent } from './pages/register/register.component';
import { ProductAddComponent } from './pages/product-add/product-add.component';



export const routes: Routes = [
  {path: '', component: LoginComponent},
  {path: 'home', component: HomeComponent},
  {path: 'meus-anuncios', component: MyAdsComponent},
  {path: 'profile', component: ProfileComponent},
  {path: 'reset-password', component: LoginResetComponent},
  {path: 'register', component: RegisterComponent},
  {path: 'product-add', component: ProductAddComponent}
];
