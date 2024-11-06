import { Component, ElementRef, ViewChild } from '@angular/core';
import { RouterLink } from '@angular/router';
import { MenuComponent } from '../../components/menu/menu.component';
import { NgFor, NgIf } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ProductService } from '../../services/api/product.service';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-product-add',
  standalone: true,
  imports: [RouterLink, MenuComponent, NgIf, NgFor, FormsModule],
  templateUrl: './product-add.component.html',
  styleUrls: ['./product-add.component.css']
})
export class ProductAddComponent {
  selectedImages: (File | null)[] = [null];
  selectedImageUrls: (string | null)[] = [null];  // URLs para exibir as imagens
  sideImages: (File | null)[] = [null, null, null];
  sideImageUrls: (string | null)[] = [null, null, null];  // URLs para exibir as imagens laterais

  categories = ['Veículos', 'Locações', 'Roupas e Calçados', 'Móveis', 'Eletrônicos', 'Outros'];
  conditions = ['Novo', 'Semi-novo', 'Em Bom Estado', 'Condições Razoáveis', 'Extensivo'];

  selectedCategory = '';
  selectedCondition = '';
  productTitle = '';
  productDescription = '';
  productPrice = '';

  @ViewChild('fileInput', { static: false }) fileInput!: ElementRef<HTMLInputElement>;

  constructor(private productService: ProductService) {}

  onFileSelected(event: Event): void {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files.length > 0) {
      const file = input.files[0];
      const reader = new FileReader();

      reader.onload = () => {
        const imageUrl = reader.result as string;

        if (!this.selectedImages[0]) {
          this.selectedImages[0] = file;
          this.selectedImageUrls[0] = imageUrl;  // Armazena a URL para exibição
        } else {
          const emptyIndex = this.sideImages.findIndex(img => img === null);
          if (emptyIndex !== -1) {
            this.sideImages[emptyIndex] = file;
            this.sideImageUrls[emptyIndex] = imageUrl;  // Armazena a URL para exibição
          }
        }
      };

      reader.readAsDataURL(file);  // Converte o arquivo para base64 para exibição
    }
  }

  formatPrice(): void {
    const numericValue = parseFloat(this.productPrice.replace(/[^\d.-]/g, ''));
    if (!isNaN(numericValue)) {
      this.productPrice = numericValue.toLocaleString('pt-BR', {
        style: 'currency',
        currency: 'BRL',
        minimumFractionDigits: 2,
      });
    }
  }

  submitProduct(): void {
    const formData = new FormData();

    const priceNumber = parseFloat(this.productPrice.replace(/[^\d.]/g, ''));
    if (isNaN(priceNumber)) {
      alert("Por favor, insira um valor válido para o preço.");
      return;
    }
    const priceValue = priceNumber.toFixed(2);

    formData.append('title', this.productTitle);
    formData.append('description', this.productDescription);
    formData.append('price', priceValue);
    formData.append('category', this.selectedCategory);
    formData.append('condition', this.selectedCondition);

    if (this.selectedImages[0]) {
      formData.append('photo1', this.selectedImages[0]);  // Envia o arquivo real
    }
    this.sideImages.forEach((file, index) => {
      if (file) {
        formData.append(`photo${index + 2}`, file);  // Envia o arquivo real
      }
    });

    this.productService.addProduct(formData).subscribe(
      (response) => {
        console.log('Produto adicionado com sucesso:', response);
      },
      (error: HttpErrorResponse) => {
        console.error('Erro ao enviar o produto:', error);
      }
    );
  }
}
