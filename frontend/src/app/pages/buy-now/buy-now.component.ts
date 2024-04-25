import { Component, OnInit } from '@angular/core';
import { SharedModule } from '../../shared/shared.module';
import { RedirectifyApiService } from '../../services/redirectify-api.service';
import { Router, RouterLink } from '@angular/router';

@Component({
  selector: 'app-buy-now',
  standalone: true,
  imports: [SharedModule,RouterLink],
  templateUrl: './buy-now.component.html',
  styleUrl: './buy-now.component.scss'
})
export class BuyNowComponent implements OnInit {

  constructor(private api:RedirectifyApiService){

  }
  plansData:any

 async ngOnInit(){
      this.plansData = await this.api.getPlans()

  }

  redirectTo(router:string){
    window.location.href = router
  }

}
