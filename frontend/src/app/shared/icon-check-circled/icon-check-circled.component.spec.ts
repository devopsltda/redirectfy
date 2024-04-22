import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IconCheckCircledComponent } from './icon-check-circled.component';

describe('IconCheckCircledComponent', () => {
  let component: IconCheckCircledComponent;
  let fixture: ComponentFixture<IconCheckCircledComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [IconCheckCircledComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(IconCheckCircledComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
