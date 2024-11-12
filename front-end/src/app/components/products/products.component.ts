import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';
import { ProductsPreviewService } from '../../services/api/product_preview.service';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-products',
  standalone: true,
  imports: [CommonModule, RouterLink, FormsModule],
  templateUrl: './products.component.html',
  styleUrls: ['./products.component.css'],
})
export class ProductsComponent implements OnInit {
  products: any[] = [];
  paginatedProducts: any[] = [];
  currentPage = 1;
  productsPerPage = 20;
  searchTerm = ''; // Variável para armazenar o termo de pesquisa

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
      },
      (error) => {
        console.error('Erro ao carregar produtos', error);
      }
    );
  }

  // Método para atualizar os produtos paginados com base na pesquisa
  updatePaginatedProducts(): void {
    const filteredProducts = this.getFilteredProducts();
    const startIndex = (this.currentPage - 1) * this.productsPerPage;
    const endIndex = startIndex + this.productsPerPage;
    this.paginatedProducts = filteredProducts.slice(startIndex, endIndex);
  }

  // Método para retornar os produtos filtrados com base no termo de pesquisa
  getFilteredProducts(): any[] {
    return this.products.filter(product =>
      product.title.toLowerCase().includes(this.searchTerm.toLowerCase())
    );
  }

  // Método para ativar a pesquisa e atualizar a lista paginada
  searchProducts(): void {
    this.currentPage = 1; // Redefine para a primeira página ao pesquisar
    this.updatePaginatedProducts();
  }

  nextPage(): void {
    if (this.currentPage * this.productsPerPage < this.getFilteredProducts().length) {
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
