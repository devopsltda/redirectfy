import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { environment } from '../../../environments/environment';


@Component({
  selector: 'app-card-subscription-plan',
  standalone: true,
  imports: [],
  templateUrl: './card-subscription-plan.component.html',
  styleUrl: './card-subscription-plan.component.scss'
})
export class CardSubscriptionPlanComponent implements OnInit {
  @Input() disable:boolean = false
  @Input() pricing:number = 555
  @Input() isAnnual:boolean = false
  @Input() planName:string = 'Nome do Plano'
  @Input() titleTextColor:string = "text-black"
  @Input() planTextColor:string = "text-black"
  @Input() bgPricingColor:string = "bg-black"
  @Input() bgCardColor:string = "bg-[#F2F2F2]"
  @Input() logoFillColor:string = "fill-black"
  @Input() typePlan!:string
  @Output() linkCheckoutEvent = new EventEmitter()
  linkCheckout!:string
  ngOnInit(): void {
      if(this.typePlan == 'Basic Mensal'){
        this.logoFillColor="fill-gray-500"
        this.bgCardColor="bg-gray-white"
        this.isAnnual=false
        this.titleTextColor="text-black-600"
        this.planTextColor="text-gray-500"
        this.linkCheckout =  environment.checkoutUrlBasicMensal
      } else if(this.typePlan == 'Pro Mensal'){
        this.logoFillColor="fill-teal-400"
        this.bgCardColor="bg-white bg-opacity-75"
        this.isAnnual=false
        this.titleTextColor="text-black text-opacity-90"
        this.planTextColor="text-teal-400"
        this.linkCheckout = environment.checkoutUrlProMensal
      } else if(this.typePlan == 'Pro Anual'){
        this.logoFillColor="fill-teal-600"
        this.bgCardColor="bg-black "
        this.isAnnual=true
        this.titleTextColor="text-white text-opacity-90"
        this.planTextColor="text-teal-600"
        this.linkCheckout = environment.checkoutUrlProAnual
      }
  }

  onClickEvent(){
    this.linkCheckoutEvent.emit(this.linkCheckout)
  }
}
