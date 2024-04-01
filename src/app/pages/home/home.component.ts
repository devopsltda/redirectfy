
import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';
import { SharedModule } from '../../shared/shared.module';
import { fadeInOutAnimation } from '../../animations/animations.module';



@Component({
  selector: 'app-home',
  standalone: true,
  imports: [SharedModule,RouterModule],
  animations:[fadeInOutAnimation],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss'
})
export class HomeComponent {

}
