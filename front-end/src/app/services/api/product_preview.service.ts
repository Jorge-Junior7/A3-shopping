import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class ProductsPreviewService {
  private baseUrl = 'http://localhost:8080/products/preview';

  constructor(private http: HttpClient) {}

  getProductsPreview(): Observable<any> {
    return this.http.get(this.baseUrl);
  }
}
