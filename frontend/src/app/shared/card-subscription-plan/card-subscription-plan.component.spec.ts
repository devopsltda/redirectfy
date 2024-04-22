import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CardSubscriptionPlanComponent } from './card-subscription-plan.component';

describe('CardSubscriptionPlanComponent', () => {
  let component: CardSubscriptionPlanComponent;
  let fixture: ComponentFixture<CardSubscriptionPlanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CardSubscriptionPlanComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(CardSubscriptionPlanComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
