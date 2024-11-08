import { Component } from '@angular/core';
import { RouterLink } from '@angular/router';
import { RegisterService } from '../../services/api/register.service';
import { NgIf } from '@angular/common';
import { OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-final-registration',
  standalone: true,
  imports: [RouterLink, NgIf],
  templateUrl: './final-registration.component.html',
  styleUrls: ['./final-registration.component.css']
})

export class FinalRegistrationComponent implements OnInit {
  recoveryPhrase: string | null = null;

  constructor(private route: ActivatedRoute) {}

  ngOnInit(): void {
    this.route.queryParams.subscribe(params => {
      this.recoveryPhrase = params['recoveryPhrase'] || null;  // Captura a frase de recuperação dos queryParams
    });
  }
}
