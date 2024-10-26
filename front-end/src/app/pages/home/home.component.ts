import { Component } from '@angular/core';
import { MenuComponent } from '../../components/menu/menu.component';
import { ProductsComponent } from '../../components/products/products.component';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [MenuComponent, ProductsComponent],  // Corrigido para ProductsComponent
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']  // Corrigido para styleUrls (plural)
})
export class HomeComponent {}
