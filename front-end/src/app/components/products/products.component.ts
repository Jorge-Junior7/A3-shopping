import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-products',
  standalone: true,
  imports: [],
  templateUrl: './products.component.html',
  styleUrl: './products.component.css'
})
export class ProductsComponent {
  @Input() product!: {
    photoCover: "";
    altPhotoCover: String;
    name: string;
    price: number;
    location: string;
  };
}

