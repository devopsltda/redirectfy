import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RedirectfyPremiumCardComponent } from './redirectfy-premium-card.component';

describe('RedirectfyPremiumCardComponent', () => {
  let component: RedirectfyPremiumCardComponent;
  let fixture: ComponentFixture<RedirectfyPremiumCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RedirectfyPremiumCardComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(RedirectfyPremiumCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
