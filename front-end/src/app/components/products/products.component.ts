import { CommonModule, NgFor } from '@angular/common';
import { Component } from '@angular/core';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-products',
  standalone: true,
  imports: [NgFor, CommonModule, RouterLink],  // Correto para usar NgFor e funcionalidades comuns do Angular
  templateUrl: './products.component.html',
  styleUrls: ['./products.component.css']  // Remover exports, não é necessário aqui
})
export class ProductsComponent {
  products = [
    {
      name: 'Gol usado',
      price: '20.000 R$',
      location: 'São Paulo',
      imageUrl: 'link-da-imagem-do-carro',
    },
    {
      name: 'Xiaomi redmi 8 semi novo',
      price: '900 R$',
      location: 'Mossoró/RN',
      imageUrl: 'link-da-imagem-do-celular',
    },
    {
      name: 'Honda 150 unico dono',
      price: '7.000 R$',
      location: 'Mossoró/RN',
      imageUrl: 'link-da-imagem-da-moto',
    },
    {
      name: 'Honda 150 unico dono',
      price: '7.000 R$',
      location: 'Mossoró/RN',
      imageUrl: 'link-da-imagem-da-moto',
    }, {
      name: 'Honda 150 unico dono',
      price: '7.000 R$',
      location: 'Mossoró/RN',
      imageUrl: 'link-da-imagem-da-moto',
    }, {
      name: 'Honda 150 unico dono',
      price: '7.000 R$',
      location: 'Mossoró/RN',
      imageUrl: 'link-da-imagem-da-moto',
    }, {
      name: 'Honda 150 unico dono',
      price: '7.000 R$',
      location: 'Mossoró/RN',
      imageUrl: 'link-da-imagem-da-moto',
    }, {
      name: 'Honda 150 unico dono',
      price: '7.000 R$',
      location: 'Mossoró/RN',
      imageUrl: 'link-da-imagem-da-moto',
    },
    // Adicione mais produtos aqui
  ];
}
