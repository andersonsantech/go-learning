package usecase

func (i *UserUsecaseImpl) Delete(id int) error {
	return i.UserRepository.Delete(id)
}
