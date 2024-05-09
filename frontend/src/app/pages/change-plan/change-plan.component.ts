import { Component, OnInit } from '@angular/core';
import { SharedModule } from '../../shared/shared.module';
import { RouterLink } from '@angular/router';
import { RedirectifyApiService } from '../../services/redirectify-api.service';

@Component({
  selector: 'app-change-plan',
  standalone: true,
  imports: [SharedModule,RouterLink],
  templateUrl: './change-plan.component.html',
  styleUrl: './change-plan.component.scss'
})
export class ChangePlanComponent implements OnInit {
  constructor(private api:RedirectifyApiService){

  }
  userData:any
  plansData:any

 async ngOnInit(){
      this.plansData = await this.api.getPlans()
      this.userData = await this.api.getUser()
      console.log(this.userData)
  }

  redirectTo(router:any){
    console.log(router)
    window.location.href = router
  }
}
