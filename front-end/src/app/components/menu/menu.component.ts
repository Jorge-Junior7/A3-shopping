import { Component, AfterViewInit } from '@angular/core';
import { RouterLink, RouterOutlet } from '@angular/router';
declare var $: any;


@Component({
  selector: 'app-menu',
  standalone: true,
  imports: [RouterLink, RouterOutlet],
  templateUrl: './menu.component.html',
  styleUrl: './menu.component.css'
})

export class MenuComponent implements AfterViewInit {

  ngAfterViewInit(): void {
    $('.toggle-menu').click(function() {
      $('.exo-menu').toggleClass('display');
    });
  }
}
