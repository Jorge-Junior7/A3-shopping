import { Component } from '@angular/core';
import { MenuComponent } from '../../components/menu/menu.component';


@Component({
  selector: 'app-my-ads',
  standalone: true,
  imports: [MenuComponent],
  templateUrl: './my-ads.component.html',
  styleUrl: './my-ads.component.css'
})
export class MyAdsComponent {

}
