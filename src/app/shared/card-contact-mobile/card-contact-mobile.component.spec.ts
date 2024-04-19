import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CardContactMobileComponent } from './card-contact-mobile.component';

describe('CardContactMobileComponent', () => {
  let component: CardContactMobileComponent;
  let fixture: ComponentFixture<CardContactMobileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CardContactMobileComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(CardContactMobileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
