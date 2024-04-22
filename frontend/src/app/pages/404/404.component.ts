import { Component } from '@angular/core';
import { SharedModule } from '../../shared/shared.module';
import { Location } from '@angular/common';
import { Router } from '@angular/router';

@Component({
  selector: 'app-404',
  standalone: true,
  imports: [SharedModule],
  templateUrl: './404.component.html',
  styleUrl: './404.component.scss'
})
export class Error404Component {

  constructor(private router:Router){}

  goBack(){
    return this.router.navigate(['/'])
  }

}
