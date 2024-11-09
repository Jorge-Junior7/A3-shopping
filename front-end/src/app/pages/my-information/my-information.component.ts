import { Component } from '@angular/core';
import { MenuComponent } from '../../components/menu/menu.component';

@Component({
  selector: 'app-my-information',
  standalone: true,
  imports: [MenuComponent],
  templateUrl: './my-information.component.html',
  styleUrl: './my-information.component.css'
})
export class MyInformationComponent {

}
