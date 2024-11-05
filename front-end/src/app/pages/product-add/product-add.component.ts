import { Component, ElementRef, ViewChild } from '@angular/core';
import { RouterLink } from '@angular/router';
import { MenuComponent } from '../../components/menu/menu.component';
import { NgFor, NgIf } from '@angular/common';

@Component({
  selector: 'app-product-add',
  standalone: true,
  imports: [RouterLink, MenuComponent, NgIf, NgFor],
  templateUrl: './product-add.component.html',
  styleUrls: ['./product-add.component.css']
})
export class ProductAddComponent {
  selectedImages: (string | ArrayBuffer | null)[] = [];
  sideImages: (string | ArrayBuffer | null)[] = [null, null, null];

  @ViewChild('fileInput', { static: false }) fileInput!: ElementRef<HTMLInputElement>;

  selectImage(): void {
    this.fileInput.nativeElement.click();
  }

  onFileSelected(event: Event): void {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files[0]) {
      const file = input.files[0];
      const reader = new FileReader();
      reader.onload = (e) => {
        if (this.selectedImages.length === 0) {
          this.selectedImages.push(e.target?.result || '');
        } else {
          const emptyIndex = this.sideImages.findIndex(img => img === null);
          if (emptyIndex !== -1) {
            this.sideImages[emptyIndex] = e.target?.result || '';
          }
        }
      };
      reader.readAsDataURL(file);
    }
  }
}
