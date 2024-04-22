import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IconMailComponent } from './icon-mail.component';

describe('IconMailComponent', () => {
  let component: IconMailComponent;
  let fixture: ComponentFixture<IconMailComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [IconMailComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(IconMailComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
