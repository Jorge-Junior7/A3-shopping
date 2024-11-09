import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class VerifyService {
  private apiUrl = 'http://localhost:8080/register';

  constructor(private http: HttpClient) {}

  verifyUserData(data: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/verify`, data);
  }

  updateUserPassword(data: any): Observable<any> {
    return this.http.post('http://localhost:8080/register/update-password', data);
  }
}
