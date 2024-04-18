import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-card-subscription-plan',
  standalone: true,
  imports: [],
  templateUrl: './card-subscription-plan.component.html',
  styleUrl: './card-subscription-plan.component.scss'
})
export class CardSubscriptionPlanComponent {
  @Input() pricing:number = 555
  @Input() isAnnual:boolean = false
  @Input() planName:string = 'Nome do Plano'
  @Input() titleTextColor:string = "text-black"
  @Input() planTextColor:string = "text-black"
  @Input() bgPricingColor:string = "bg-black"
  @Input() bgCardColor:string = "bg-[#F2F2F2]"
  @Input() logoFillColor:string = "fill-black"
}
