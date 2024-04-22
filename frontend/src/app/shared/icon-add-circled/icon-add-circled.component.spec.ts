import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IconAddCircledComponent } from './icon-add-circled.component';

describe('IconAddCircledComponent', () => {
  let component: IconAddCircledComponent;
  let fixture: ComponentFixture<IconAddCircledComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [IconAddCircledComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(IconAddCircledComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
