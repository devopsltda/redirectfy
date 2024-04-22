import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IconVisaCardComponent } from './icon-visa-card.component';

describe('IconVisaCardComponent', () => {
  let component: IconVisaCardComponent;
  let fixture: ComponentFixture<IconVisaCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [IconVisaCardComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(IconVisaCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
