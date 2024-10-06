import { Routes } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';
import { HomeComponent } from './pages/home/home.component';
import { MyAdsComponent } from './pages/my-ads/my-ads.component';
import { ProfileComponent } from './pages/profile/profile.component';

export const routes: Routes = [
  {path: '', component: LoginComponent},
  {path: 'home', component: HomeComponent},
  {path: 'myAds', component: MyAdsComponent},
  {path: 'profile', component: ProfileComponent},
];
