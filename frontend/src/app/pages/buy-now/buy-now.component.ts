import { Component } from '@angular/core';
import { SharedModule } from '../../shared/shared.module';

@Component({
  selector: 'app-buy-now',
  standalone: true,
  imports: [SharedModule],
  templateUrl: './buy-now.component.html',
  styleUrl: './buy-now.component.scss'
})
export class BuyNowComponent {

}
