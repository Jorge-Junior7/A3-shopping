import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';
import { ProductsPreviewService } from '../../services/api/product_preview.service';

@Component({
  selector: 'app-products',
  standalone: true,
  imports: [CommonModule, RouterLink],
  templateUrl: './products.component.html',
  styleUrls: ['./products.component.css'],
})
export class ProductsComponent implements OnInit {
  products: any[] = [];
  paginatedProducts: any[] = [];
  currentPage = 1;
  productsPerPage = 20;

  constructor(private productsPreviewService: ProductsPreviewService) {}

  ngOnInit(): void {
    this.loadProducts();
  }

  loadProducts(): void {
    this.productsPreviewService.getProductsPreview().subscribe(
      (data) => {
        this.products = data;
        this.updatePaginatedProducts();
        console.log('Produtos carregados:', this.products);
        this.products.forEach((product) => {
          console.log('Imagem do produto:', product.photo1);
          console.log('Preço do produto:', product.price);
        });
      },
      (error) => {
        console.error('Erro ao carregar produtos', error);
      }
    );
  }

  updatePaginatedProducts(): void {
    const startIndex = (this.currentPage - 1) * this.productsPerPage;
    const endIndex = startIndex + this.productsPerPage;
    this.paginatedProducts = this.products.slice(startIndex, endIndex);
  }

  nextPage(): void {
    if (this.currentPage * this.productsPerPage < this.products.length) {
      this.currentPage++;
      this.updatePaginatedProducts();
    }
  }

  previousPage(): void {
    if (this.currentPage > 1) {
      this.currentPage--;
      this.updatePaginatedProducts();
    }
  }
}
