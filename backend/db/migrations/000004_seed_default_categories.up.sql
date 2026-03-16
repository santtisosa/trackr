-- Create function to insert default categories for a new user
CREATE OR REPLACE FUNCTION public.insert_default_categories(user_id UUID)
RETURNS VOID AS $$
BEGIN
    INSERT INTO public.categories (user_id, name, icon, color, is_default) VALUES
        (user_id, 'Alimentación',    'shopping-cart', '#FF6B6B', TRUE),
        (user_id, 'Transporte',      'car',           '#4ECDC4', TRUE),
        (user_id, 'Vivienda',        'home',          '#45B7D1', TRUE),
        (user_id, 'Salud',           'heart-pulse',   '#96CEB4', TRUE),
        (user_id, 'Educación',       'book-open',     '#FFEAA7', TRUE),
        (user_id, 'Entretenimiento', 'gamepad-2',     '#DDA0DD', TRUE),
        (user_id, 'Ropa',            'shirt',         '#98D8C8', TRUE),
        (user_id, 'Otros',           'package',       '#B0B0B0', TRUE);
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;

-- Update handle_new_user trigger to also seed default categories
CREATE OR REPLACE FUNCTION public.handle_new_user()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO public.profiles (id, email, full_name)
    VALUES (NEW.id, NEW.email, NEW.raw_user_meta_data->>'full_name');

    PERFORM public.insert_default_categories(NEW.id);

    RETURN NEW;
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;
