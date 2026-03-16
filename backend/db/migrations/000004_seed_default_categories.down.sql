-- Drop insert_default_categories function
DROP FUNCTION IF EXISTS public.insert_default_categories(UUID);

-- Revert handle_new_user to previous version (without seeding categories)
CREATE OR REPLACE FUNCTION public.handle_new_user()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO public.profiles (id, email, full_name)
    VALUES (NEW.id, NEW.email, NEW.raw_user_meta_data->>'full_name');

    RETURN NEW;
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;
