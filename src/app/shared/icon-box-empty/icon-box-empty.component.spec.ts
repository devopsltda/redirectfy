import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IconBoxEmptyComponent } from './icon-box-empty.component';

describe('IconBoxEmptyComponent', () => {
  let component: IconBoxEmptyComponent;
  let fixture: ComponentFixture<IconBoxEmptyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [IconBoxEmptyComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(IconBoxEmptyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
