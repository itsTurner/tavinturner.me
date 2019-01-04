class ApplicationController < ActionController::Base
  protect_from_forgery with: :exception

  module ApplicationHelper
    def title(text)
      content_for :title, text
    end
  end
end
